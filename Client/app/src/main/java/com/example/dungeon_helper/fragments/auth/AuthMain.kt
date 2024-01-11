package com.example.dungeon_helper.fragments.auth

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAuthMainBinding
import com.example.dungeon_helper.databinding.FragmentAuthRegistrationBinding

class AuthMain : Fragment() {

    companion object {
        fun newInstance() = AuthMain()
    }

    private lateinit var viewModel: AuthMainViewModel

    private var _binding: FragmentAuthMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authMainViewModel = ViewModelProvider(this)[AuthMainViewModel::class.java]

        _binding = FragmentAuthMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textAuthMain
        authMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root

    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val regBtn = binding.registrationBtn
        val restoreBtn = binding.restorePwdBtn
        val loginBtn = binding.loginBtn
        val loginGoogleBtn = binding.loginGoogleBtn

        regBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_auth_to_authRegistration)
        }

        restoreBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_auth_to_authRestorePwd)
        }
    }

}