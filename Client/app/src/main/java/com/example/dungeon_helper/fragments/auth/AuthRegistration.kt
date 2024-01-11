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
import com.example.dungeon_helper.databinding.FragmentAuthRegistrationBinding
import com.example.dungeon_helper.databinding.FragmentAuthRestorePwdBinding

class AuthRegistration : Fragment() {

    companion object {
        fun newInstance() = AuthRegistration()
    }

    private lateinit var viewModel: AuthRegistrationViewModel

    private var _binding: FragmentAuthRegistrationBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRegistrationViewModel = ViewModelProvider(this)[AuthRegistrationViewModel::class.java]

        _binding = FragmentAuthRegistrationBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textRegistration
        authRegistrationViewModel.text.observe(viewLifecycleOwner) {
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
        val backBtn = binding.backBtn
        val regBtn = binding.registrationBtn

        backBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_authRegistration_to_auth)
        }

    }

}