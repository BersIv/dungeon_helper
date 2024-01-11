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
import com.example.dungeon_helper.databinding.FragmentAuthRestorePwdBinding
import com.example.dungeon_helper.databinding.FragmentCharacterMainBinding

class AuthRestorePwd : Fragment() {

    companion object {
        fun newInstance() = AuthRestorePwd()
    }

    private lateinit var viewModel: AuthRestorePwdViewModel

    private var _binding: FragmentAuthRestorePwdBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRestorePwdViewModel = ViewModelProvider(this)[AuthRestorePwdViewModel::class.java]

        _binding = FragmentAuthRestorePwdBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textRestorePwd
        authRestorePwdViewModel.text.observe(viewLifecycleOwner) {
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
        val restoreBtn = binding.restoreBtn

        backBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd_to_auth)
        }

    }
}