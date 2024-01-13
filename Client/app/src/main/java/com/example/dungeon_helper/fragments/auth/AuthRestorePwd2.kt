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
import com.example.dungeon_helper.databinding.FragmentAuthRestorePwd2Binding
class AuthRestorePwd2 : Fragment() {

    companion object {
        fun newInstance() = AuthRestorePwd2()
    }

    private lateinit var viewModel: AuthRestorePwd2ViewModel

    private var _binding: FragmentAuthRestorePwd2Binding? = null
    private  val binding get() = _binding!!
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val authRestorePwd2ViewModel = ViewModelProvider(this)[AuthRestorePwd2ViewModel::class.java]
        _binding = FragmentAuthRestorePwd2Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textRestorePwd2
        authRestorePwd2ViewModel.text.observe(viewLifecycleOwner)
        {
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
        val backBtn2 = binding.backBtn2
        val savePwdBtn = binding.savePwdBtn

        backBtn2.setOnClickListener{
            (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd2_to_auth)
        }
        savePwdBtn.setOnClickListener {
            (activity as AuthActivity).navController.navigate(R.id.action_authRestorePwd2_to_auth)
        }
    }

}